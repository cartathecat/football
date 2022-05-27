// v3.0
const DEBUG = true;

/*
Loaded when the neo4jsearchresults is loaded
*/
function init() {
    console.log(window.location.pathname)
  
    if (window.location.pathname == '/index') displayCorrectPageHandler()
  
}

function displayCorrectPageHandler() {

    console.log("Index : displayCorrectPageHandler");
    var htmlString = ""

    if(/Android|webOS|iPhone|iPad|iPod|BlackBerry|IEMobile|Opera Mini/i.test(navigator.userAgent)){
        // true for mobile device
        console.log("mobile")
        htmlString += '<div>mobile</div><div id="scoreaxis-widget-41174" style="border-width:1px;border-color:rgba(0, 0, 0, 0.15);border-style:solid;border-radius:8px;padding:10px;background:rgb(255, 255, 255);width:auto;height:135%" data-reactroot=""><iframe id="Iframe" src="https://www.scoreaxis.com/widget/standings-widget/8?widgetRows=1%2C1%2C1%2C1%2C1%2C1%2C1%2C1%2C0%2C1&amp;removeBorders=0&amp;widgetHomeAwayTabs=0&amp;inst=41174" style="width:100%;height:100%;border:none;transition:all 300ms ease"></iframe><script>window.addEventListener("DOMContentLoaded",event=>{window.addEventListener("message",event=>{if(event.data.appHeight&&"41174"==event.data.inst){let container=document.querySelector("#scoreaxis-widget-41174 iframe");container&&(container.style.height=parseInt(event.data.appHeight)+"px")}},!1)});</script></div>'

    } else{
        // false for not mobile device
        console.log("desktop")
        htmlString += '<div>desktop</div><div id="scoreaxis-widget-9b29b-desktop" style="border-width:1px;border-color:rgba(0, 0, 0, 0.15);border-style:solid;border-radius:8px;padding:10px;background:rgb(255, 255, 255);width:auto;height:135%;" data-reactroot=""><iframe id="Iframe-desktop" title="Iframe-desktop" src="https://www.scoreaxis.com/widget/standings-widget/8?&amp;inst=9b29b" style="width:100%;height:100%;border:none;transition:all 300ms ease"></iframe><script>window.addEventListener("DOMContentLoaded",event=>{window.addEventListener("message",event=>{if(event.data.appHeight&&"9b29b"==event.data.inst){let container=document.querySelector("#scoreaxis-widget-9b29b iframe");container&&(container.style.height=parseInt(event.data.appHeight)+"px")}},!1)});</script></div>'
    }
    document.getElementById("leagueTable").innerHTML = htmlString;
}
