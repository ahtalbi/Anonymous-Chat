import { API_URL } from "../../../config.js";

async function SetHtml(app, base, page) {
    try {
        let res = await fetch(`${base + page}/${page}.html`);
        res = await res.text();
        console.log(res);
        
        app.innerHTML = "";
        app.innerHTML = res;
    } catch (error) {
        console.log(error);
    }
}

export function RenderPage(app, page) {
    let base = API_URL + "/src/pages/";
    
    SetHtml(app, base, page);

}