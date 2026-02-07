import { Router } from "../../packages/router.js";
import { RenderPage } from "./utils/render.js";

export function initServer() {
    let app = document.getElementById("app");
    let router = new Router();
    let OnError404 = () => { app.innerHTML = "<h1>404</h1>" };
    let routes = {
        "/": {auth: true, handler: () => RenderPage(app, "messages")},
        "/login": {auth: false, handler: () => RenderPage(app, "login")}
    };
    
    for (let route in routes) {
        router.on(route, routes[route].handler)
    }
    router.listen(OnError404);

    let path = window.location.pathname;
    
    router.navigate(path);
    // else if (routes[path].auth && )
}