import { EventsManager } from "../../../packages/event.js";

export let EventManager = null;
export function initGlobalEventManager() {
    EventManager = {
        submit: new EventsManager("submit"),
        click: new EventsManager("click"),
    };
}