use crate::error::CustomError;
use actix_web::{get, web, HttpResponse, Responder, Result, Scope};
use graphql_client::{reqwest::post_graphql, GraphQLQuery};
use reqwest::Client;
use serde_json::{json, Value};

pub fn api() -> Scope {
    web::scope("/api")
        .service(get_contributions)
        .service(get_projects)
        .service(get_resume)
}

// this is the type that graphql_client needs from the contributions query
type URI = String;
// this is the type that graphql_client needs from the repos query
type Date = String;

#[derive(GraphQLQuery)]
#[graphql(
    schema_path = "src/github/schema.docs.graphql",
    query_path = "src/github/contributions.graphql",
    response_derives = "Debug,Serialize"
)]
struct Contributions;

#[derive(GraphQLQuery)]
#[graphql(
    schema_path = "src/github/schema.docs.graphql",
    query_path = "src/github/repos.graphql",
    response_derives = "Debug,Serialize"
)]
struct Repos;

#[get("/contributions")]
async fn get_contributions(client: web::Data<Client>) -> Result<impl Responder, CustomError> {
    // post GraphQL query to GitHub API
    let variables = contributions::Variables {
        login: "waltzofpearls".to_string(),
    };
    let response =
        post_graphql::<Contributions, _>(&client, "https://api.github.com/graphql", variables)
            .await?;

    // check for GraphQL errors
    if let Some(errors) = response.errors {
        let err = json!({
            "code": 400,
            "message": "GraphQL errors",
            "details": errors
        });
        return Ok(HttpResponse::BadRequest().json(err));
    }

    // get the data attribute from the response
    let data: contributions::ResponseData = response
        .data
        .ok_or(CustomError::from(String::from("missing response data")))?;

    // convert the data to JSON
    let json_response = serde_json::to_value(&data)?;

    // return the JSON response
    Ok(HttpResponse::Ok().json(json_response))
}

#[get("/projects")]
async fn get_projects(client: web::Data<Client>) -> Result<impl Responder, CustomError> {
    // post GraphQL query to GitHub API
    let variables = repos::Variables {
        login: "waltzofpearls".to_string(),
    };
    let response =
        post_graphql::<Repos, _>(&client, "https://api.github.com/graphql", variables).await?;

    // check for GraphQL errors
    if let Some(errors) = response.errors {
        let err = json!({
            "code": 400,
            "message": "GraphQL errors",
            "details": errors
        });
        return Ok(HttpResponse::BadRequest().json(err));
    }

    // get the data attribute from the response
    let data: repos::ResponseData = response
        .data
        .ok_or(CustomError::from(String::from("missing response data")))?;

    // convert the data to JSON
    let json_response = serde_json::to_value(&data)?;

    // return the JSON response
    Ok(HttpResponse::Ok().json(json_response))
}

#[get("/resume")]
async fn get_resume() -> Result<impl Responder, CustomError> {
    // read a json file and return the contents as json
    let json_response: Value = serde_json::from_str(include_str!("./resume/doc.json"))
        .map_err(|err| CustomError::from(format!("Failed to read JSON: {}", err)))?;

    // return the JSON response
    Ok(HttpResponse::Ok().json(json_response))
}
