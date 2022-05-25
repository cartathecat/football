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

    if(/Android|webOS|iPhone|iPad|iPod|BlackBerry|IEMobile|Opera Mini/i.test(navigator.userAgent)){
        // true for mobile device
        console.log("mobile")
   //     document.getElementById("mobile").style.visibility = "visible";
   //     document.getElementById("desktop").style.visibility = "hidden";

      }else{
        // false for not mobile device
        console.log("desktop")
   //     document.getElementById("mobile").style.visibility = "hidden";
   //     document.getElementById("desktop").style.visibility = "visible";
      }

}
