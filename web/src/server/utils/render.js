import { API_URL } from "../../../config.js";

async function SetCss(base, page) {
    const href = `${base}${page}/${page}.css`;
    if (!(await fileExists(href, "css"))) return;
    document.querySelectorAll('link[data-page-css]').forEach(l => l.remove());
    const link = document.createElement("link");
    link.rel = "stylesheet";
    link.href = href;
    link.dataset.pageCss = page;
    document.head.appendChild(link);
    await new Promise((r) => {
        link.onload = r;
    });
}

async function SetJavaScript(base, page) {
    let src = `${base}${page}/${page}.js`;
    if (!(await fileExists(src, "js"))) return;
    document.querySelectorAll('script[data-page-script]').forEach(s => s.remove());
    let script = document.createElement("script");
    script.type = "module";
    script.src = src;
    script.dataset.pageScript = page;
    document.body.appendChild(script);
    await new Promise((resolve, reject) => {
        script.onload = resolve;
    });
}

async function SetHtml(app, base, page) {
    let res = await fetch(`${base + page}/${page}.html`);
    res = await res.text();
    app.innerHTML = res;
}

export async function RenderPage(app, page) {
    let base = API_URL + "/src/pages/";

    await SetHtml(app, base, page);
    await SetJavaScript(base, page);
    await SetCss(base, page);
}

async function fileExists(url, type) {
    const res = await fetch(url, { method: "HEAD" });
    const ct = (res.headers.get("content-type") || "").toLowerCase();
    switch (type) {
        case "js":
            return ct.includes("javascript");

        case "css":
            return ct.includes("text/css");
    }
    return res.ok;
}