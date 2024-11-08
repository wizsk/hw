// main.js. Don't remvoe this comment!
const inputClearBtn = document.getElementById("inputClear")

inputClearBtn.addEventListener('click', () => {
  input.value = "";
})

input.addEventListener("keydown", function (e) {
  if (e.key === "Enter") {
    setSubmiter("root");
  }
});

function setSubmiter(str) {
  if (input.value === "") {
    console.log("input valule is <empty>");
    return;
  }

  let href = "";
  if (str === "root") {
    href = `/r?w=${input.value}`;
  } else if (str === "text") {
    href = `/t?w=${input.value}`;
  } else {
    console.log(`${srt} aita ki vhai?`);
    return;
  }
  loadPage(href);
}

function loadPage(href) {
  const a = document.createElement("a");
  a.style.display = "hidden";
  a.href = href;
  document.body.append(a);
  a.click();
}


function fontSizePlus() {
  const curr = parseFloat(window.getComputedStyle(document.body).fontSize);
  document.body.style.fontSize = `${curr + 2}px`;
  document.body.style.lineHeight = '';
}


function fontSizeMinus() {
  const curre = parseFloat(window.getComputedStyle(document.body).fontSize);
  document.body.style.fontSize = `${curre - 2}px`
  document.body.style.lineHeight = '';
}

let currentTheme = "light";
const availabeThemes = ["light", "reading", "dark"];
const mobileHamTouggle = document.getElementById("mobileHamTouggle");
const mobileHamTouggleDot = document.getElementById("mobileHamTouggleDot");
const mobileHamTouggleCross = document.getElementById("mobileHamTouggleCross");
const mobileHamMenu = document.getElementById("mobileHamMenu");
let isMobileHamShowing = false;

mobileHamTouggle.addEventListener('click', () => {
  console.log(mobileHamMenu.classLis)
  if (isMobileHamShowing) {
    mobileHamMenu.classList.add("hidden");
    mobileHamTouggleCross.classList.add("hidden");
    mobileHamTouggleDot.classList.remove("hidden");
    isMobileHamShowing = false;
  } else {
    mobileHamMenu.classList.remove("hidden");
    mobileHamTouggleDot.classList.add("hidden");
    mobileHamTouggleCross.classList.remove("hidden");
    isMobileHamShowing = true;
  }
})

window.onload = () => {
  document.getElementById(currentTheme).classList.remove('hidden');
}

function setTheme() {
  for (let i = 0; i < availabeThemes.length; i++) {
    document.getElementById(availabeThemes[i]).classList.add('hidden');
  }
  document.getElementById(currentTheme).classList.remove('hidden');
}

document.getElementById("themeTouggle").addEventListener('click', () => {
  console.log("what");
  let next = availabeThemes.indexOf(currentTheme) + 1;
  if (next === availabeThemes.length) {
    next = 0;
  }
  currentTheme = availabeThemes[next];
  setTheme();
})

/*
// nav hide
let lastScrollTop = 0;
const navbar = document.getElementById("navbar");

window.addEventListener("scroll", function () {
  let currentScroll = window.pageYOffset || document.documentElement.scrollTop;

  if (currentScroll > lastScrollTop) {
    // Scroll Down - Hide Navbar
    if (this.window.innerWidth <= 640) {
      navbar.style.transform = "translateY(100%)";
    } else {
      navbar.style.transform = "translateY(-100%)";
    }
  } else {
    // Scroll Up - Show Navbar
    navbar.style.transform = "translateY(0)";
  }

  lastScrollTop = currentScroll <= 0 ? 0 : currentScroll; // Prevent negative scroll
});
*/