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
        navigation.addEventListener("navigate", (e) => {
            e.intercept({
                handler() {
                    const url = new URL(e.destination.url);
                    const path = url.pathname;
                    console.log(url, path);
                    console.log(this.#Routes);
                    
                    if (!this.#Routes[path]) { OnError404(); return }
                    console.log(this.#Routes[path]);
                    
                    this.#Routes[path]();
                }
            })
        });
        return this;
    }
}