var body = document.getElementsByTagName("body")[0];
var datetimeElement = document.getElementById("datetime");
var toolbarElement = document.getElementById("toolbar");

// updates the clock and date in the taskbar
function updateClock() {
  var now = new Date(), // current date
    months = [
      "Jan",
      "Feb",
      "Mar",
      "Apr",
      "May",
      "Jun",
      "Jul",
      "Aug",
      "Sep",
      "Nov",
      "Dec",
    ];
  if (now.getSeconds() < 10) {
    if (now.getMinutes() < 10) {
      var time =
        now.getHours() + ":0" + now.getMinutes() + ":0" + now.getSeconds();
    } else {
      var time =
        now.getHours() + ":" + now.getMinutes() + ":0" + now.getSeconds();
    }
  } else {
    if (now.getMinutes() < 10) {
      var time =
        now.getHours() + ":0" + now.getMinutes() + ":" + now.getSeconds();
    } else {
      var time =
        now.getHours() + ":" + now.getMinutes() + ":" + now.getSeconds();
    }
  }
  var date = [now.getDate(), months[now.getMonth()], now.getFullYear()].join(
    " ",
  );
  datetimeElement.innerHTML = time + " | " + date;
  // update every second
  setTimeout(updateClock, 100);
}

// updates whether the taskbar is visible
// function updateTaskbar(e) {
//   var mouseY = e.clientY;
//   var height = window.innerHeight;

//   if (height - mouseY <= 50) {
//     toolbarElement.transition = 0;
//     toolbarElement.style.height = "40px";
//     toolbarElement.transition = "600ms";
//     toolbarElement.style.bottom = "0px";
//   } else {
//     toolbarElement.transition = "600ms";
//     toolbarElement.style.bottom = "-40px";
//     toolbarElement.transition = 0;
//     toolbarElement.style.height = "0px";
//   }
// }

// checks whether it is the login screen, and determines whether to show the context menu
// if (document.addEventListener && !("login" in document.title.toLowerCase)) {
//   document.addEventListener(
//     "contextmenu",
//     function (e) {
//       e.preventDefault();
//       var contextMenu = document.getElementById("context-menu");
//       contextMenu.style.display = "block";
//       contextMenu.style.top = String(e.clientY) + "px";
//       contextMenu.style.left = String(e.clientX) + "px";
//     },
//     false,
//   );
// }

// // closes the context menu onclick
// function closeContextMenu() {
//   var contextMenu = document.getElementById("context-menu");
//   contextMenu.style.display = "none";
// }

// // toggles hiding taskbar
// function taskbar() {
//   if (document.getElementById("taskbarToggle").innerHTML.includes("✓")) {
//     document
//       .getElementsByTagName("html")[0]
//       .removeEventListener("mousemove", updateTaskbar);
//     toolbarElement.style.bottom = "0px";
//     toolbarElement.style.height = "40px";
//     document.getElementById("taskbarToggle").innerHTML =
//       "Close Taskbar Automatically 𐄂";
//   } else if (document.getElementById("taskbarToggle").innerHTML.includes("𐄂")) {
//     document
//       .getElementsByTagName("html")[0]
//       .addEventListener("mousemove", updateTaskbar);
//     toolbarElement.style.bottom = "-40px";
//     toolbarElement.style.height = "0px";
//     document.getElementById("taskbarToggle").innerHTML =
//       "Close Taskbar Automatically ✓";
//   }
// }

function detectOS() {
  let userAgent = window.navigator.userAgent,
    platform = window.navigator.platform,
    macosPlatforms = ["Macintosh", "MacIntel", "MacPPC", "Mac68K"],
    windowsPlatforms = ["Win32", "Win64", "Windows", "WinCE"],
    iosPlatforms = ["iPhone", "iPad", "iPod"],
    os = null;

  if (macosPlatforms.indexOf(platform) != -1) {
    os = "Mac OS";
  } else if (iosPlatforms.indexOf(platform) != -1) {
    os = "iOS";
  } else if (windowsPlatforms.indexOf(platform) != -1) {
    os = "Windows";
  } else if (/Android/.test(userAgent)) {
    os = "Android";
  } else if (!os && /Linux/.test(platform)) {
    os = "Linux";
  }

  return os;
}

function search() {
  // do stuff here
  console.log("search");
  return;
}
function pageCheck(e) {
  var npages = document.getElementsByClassName("page").length;
  var pages = document.getElementsByClassName("page");
  var curr_page = document
    .getElementById("app-container")
    .getAttribute("curr-page");
  if (e.which == 39) {
    if (npages > curr_page) {
      for (let i = 0; i < npages; i++) {
        pages[i].style.transform =
          `translateX(${parseInt(curr_page) * -100}vw)`;
      }
      document
        .getElementById("app-container")
        .setAttribute("curr-page", parseInt(curr_page) + 1);
    }
  } else if (e.which == 37) {
    if (curr_page > 1) {
      for (let i = 0; i < npages; i++) {
        pages[i].style.transform =
          `translateX(${parseInt((parseInt(curr_page) - 2) * -100)}vw)`;
      }
      document
        .getElementById("app-container")
        .setAttribute("curr-page", parseInt(curr_page) - 1);
    }
  }
}

function changePage(step) {
  var npages = document.getElementsByClassName("page").length;
  var pages = document.getElementsByClassName("page");
  var curr_page = document
    .getElementById("app-container")
    .getAttribute("curr-page");
  if (step == 1) {
    if (npages > curr_page) {
      for (let i = 0; i < npages; i++) {
        pages[i].style.transform =
          `translateX(${parseInt(curr_page) * -100}vw)`;
      }
      document
        .getElementById("app-container")
        .setAttribute("curr-page", parseInt(curr_page) + 1);
    }
  } else if (step == -1) {
    if (curr_page > 1) {
      for (let i = 0; i < npages; i++) {
        pages[i].style.transform =
          `translateX(${parseInt((parseInt(curr_page) - 2) * -100)}vw)`;
      }
      document
        .getElementById("app-container")
        .setAttribute("curr-page", parseInt(curr_page) - 1);
    }
  }
  if (parseInt(curr_page) + step == npages) {
    document.getElementById("right").style.display = "none";
  } else if (parseInt(curr_page) + step == 1) {
    document.getElementById("left").style.display = "none";
  } else {
    document.getElementById("right").style.display = "block";
    document.getElementById("left").style.display = "block";
  }
}
var OSName = detectOS();
document.getElementById("OS").innerHTML = "System is: amd64/macos"; //+ navigator.platform;
document.getElementById("search").addEventListener("keyup", (e) => {
  if (e.key === "Enter" || e.keyCode == 13) {
    search();
  }
});
// adds the event listener to the html element so that it can be removed later
// document
//   .getElementsByTagName("html")[0]
//   .addEventListener("mousemove", updateTaskbar);
updateClock(); // initial call for running the clock
