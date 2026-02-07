import { API_URL } from "../../config.js";
import { Router } from "../../packages/router.js";
import { RenderPage } from "./utils/render.js";

export let ClientRouter = new Router();
export async function initServer() {
    let app = document.getElementById("app");
    let OnError404 = () => { app.innerHTML = "<h1>404</h1>" };
    let routes = {
        "/": { auth: true, handler: () => RenderPage(app, "messages") },
        "/login": { auth: false, handler: () => RenderPage(app, "login") }
    };

    for (let route in routes) {
        ClientRouter.on(route, routes[route].handler)
    }
    ClientRouter.listen(OnError404);

    let path = window.location.pathname;
    try {
        let res = await fetch(API_URL + "/api/session")
        if (!res.ok && routes[path].auth) ClientRouter.navigate("/login");
        else ClientRouter.navigate(path);
    } catch (error) { console.log(error) };
}