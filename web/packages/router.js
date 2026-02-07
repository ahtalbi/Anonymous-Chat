export class Router {
    #Routes = Object.create(null);

    on(route, handler) {
        this.#Routes[route] = handler;
        return this;
    }

    navigate(route) {
        navigation.navigate(route);
        return this;
    }

    listen(OnError404) {
        const router = this;
        navigation.addEventListener("navigate", (e) => {
            e.intercept({
                handler() {
                    const path = new URL(e.destination.url).pathname;

                    if (!router.#Routes[path]) {
                        OnError404();
                        return;
                    }

                    router.#Routes[path]();
                }
            });
        });
        return this;
    }
}