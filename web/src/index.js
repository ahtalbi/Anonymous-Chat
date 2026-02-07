import { initServer } from "./server/server.js";
import { initGlobalEventManager } from "./server/utils/events.js";

function main() {
    initGlobalEventManager();
    initServer();
}

main();