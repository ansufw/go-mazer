// Don't forget to load bootstrap js in your project
// Here, we load it in our app.js
import bootstrapBundle from "bootstrap/dist/js/bootstrap.bundle"
window.bootstrap = bootstrapBundle

import PerfectScrollbar from "perfect-scrollbar"
window.PerfectScrollbar = PerfectScrollbar


// We could import PerfectScrollbar directly in the sidebar module
import "./../static/js/components/sidebar"

console.log("debug >>>>>> mazer.js loaded delete meeee delme")

