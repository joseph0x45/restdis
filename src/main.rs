use std::net::SocketAddr;

use axum::{response::Html, Router, routing::get};

#[tokio::main]
async fn main() {
    let routes = Router::new().route("/", get(|| async {Html("Hello")}));
    let addr = SocketAddr::from(([127, 0, 0, 1], 8080));
    axum::Server::bind(&addr).serve(routes.into_make_service()).await.unwrap()
}
