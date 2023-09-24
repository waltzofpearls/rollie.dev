pub mod error;
pub mod routes;

use actix_files::Files;
use actix_web::web::Data;
use actix_web::{middleware::Logger, App, HttpServer};
use log::info;
use reqwest::Client;
use std::env;

#[actix_web::main]
async fn main() -> std::io::Result<()> {
    env_logger::init_from_env(env_logger::Env::new().default_filter_or("info"));
    info!("Starting up...");

    let listen_http = env::var("LISTEN_HTTP").unwrap_or_else(|_| "0.0.0.0:3000".to_string());
    info!("Listening on HTTP address: {}", listen_http);

    let github_token = std::env::var("GITHUB_TOKEN").expect("Missing GITHUB_TOKEN env var");
    let client = Client::builder()
        .user_agent("graphql-rust/0.10.0")
        .default_headers(
            std::iter::once((
                reqwest::header::AUTHORIZATION,
                reqwest::header::HeaderValue::from_str(&format!("Bearer {}", github_token))
                    .unwrap(),
            ))
            .collect(),
        )
        .build()
        .map_err(|err| std::io::Error::new(std::io::ErrorKind::Other, err.to_string()))?;

    HttpServer::new(move || {
        App::new()
            .wrap(Logger::default())
            .app_data(Data::new(client.clone()))
            .service(routes::api())
            .service(Files::new("/", "./static/").index_file("index.html"))
    })
    .bind(listen_http)?
    .run()
    .await
}
