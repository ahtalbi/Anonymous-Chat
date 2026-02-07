import { API_URL } from "../../../config.js";
import { ClientRouter } from "../../server/server.js";
import { EventManager } from "../../server/utils/events.js";

EventManager.submit.on("LoginForm", async (ele) => {
    let payload = {
        nickname: ele.nickname.value.trim()
    };
    try {
        let res = await fetch(API_URL + "/api/auth", {
            method: "POST",
            headers: { "Content-Type": "application/json" },
            credentials: "include",
            body: JSON.stringify(payload)
        })
        if (!res.ok) {
            const msg = await res.text();
            throw new Error(msg || `HTTP ${res.status}`);
        }
        ClientRouter.navigate("/");
        console.log(res);
    } catch (error) {
        console.log(error);
    }
});