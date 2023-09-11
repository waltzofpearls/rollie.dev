use actix_web::http::StatusCode;
use actix_web::{error, HttpResponse};
use log::info;
use serde_json::json;

#[derive(Debug)]
pub enum CustomErrorType {
    ReqwestError,
    UserError(String),
}

#[derive(Debug)]
pub struct CustomError {
    pub message: Option<String>,
    pub err_type: CustomErrorType,
}
impl CustomError {
    pub fn message(&self) -> String {
        match &self.message {
            Some(c) => c.clone(),
            None => String::from(""),
        }
    }
}

impl std::fmt::Display for CustomError {
    fn fmt(&self, f: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
        write!(f, "{:?}", self)
    }
}

impl From<reqwest::Error> for CustomError {
    fn from(err: reqwest::Error) -> CustomError {
        CustomError {
            message: Some(err.to_string()),
            err_type: CustomErrorType::ReqwestError,
        }
    }
}

impl From<serde_json::Error> for CustomError {
    fn from(err: serde_json::Error) -> CustomError {
        CustomError {
            message: Some(err.to_string()),
            err_type: CustomErrorType::ReqwestError,
        }
    }
}

impl From<String> for CustomError {
    fn from(err: String) -> CustomError {
        CustomError {
            message: Some(err.clone()),
            err_type: CustomErrorType::UserError(err),
        }
    }
}

impl error::ResponseError for CustomError {
    fn status_code(&self) -> StatusCode {
        match self.err_type {
            CustomErrorType::ReqwestError => StatusCode::INTERNAL_SERVER_ERROR,
            CustomErrorType::UserError(_) => StatusCode::INTERNAL_SERVER_ERROR,
        }
    }
    fn error_response(&self) -> HttpResponse {
        info!("error: {:?}", self);
        let err = json!({
            "code": self.status_code().as_u16(),
            "message": self.message()
        });
        HttpResponse::build(self.status_code()).json(err)
    }
}
