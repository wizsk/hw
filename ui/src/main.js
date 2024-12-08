// main.js. Don't remvoe this comment!
const inputClearBtn = document.getElementById("inputClear")
const input = document.getElementById("input");
const inputSugg = document.getElementById("inputSuggestions");

const mainSection = document.querySelector("main");

console.log(mainSection);
let suggHinnedTimeoutId;

function isSuggHidden() {
  return inputSugg.classList.contains('hidden');
}

window.addEventListener('scroll', function (e) {
  if (!isSuggHidden()) setTimeout(() => inputSugg.classList.add('hidden'), 100);
})


document.addEventListener('click', (e) => {
  console.log(e.target);
  if (!inputSugg.contains(e.target) && !input.contains(e.target)) {
    if (!isSuggHidden()) inputSugg.classList.add('hidden');
  }
});

document.addEventListener('scroll', (e) => {
  if (!isSuggHidden()) inputSugg.classList.add('hidden');
});


let inputTimeout;
input.addEventListener('input', () => {
  clearTimeout(inputTimeout);

  inputTimeout = setTimeout(async () => {
    const res = await fetch(`/sugg?w=${encodeURIComponent(input.value)}`, { method: "POST" });
    if (res.ok) {
      const data = await res.text();
      if (data) {
        inputSugg.innerHTML = data;
        inputSugg.classList.remove("hidden");
      }
    }
  }, 400);
})

document.addEventListener('keydown', (e) => {
  if (document.activeElement === input) {
    if (e.code === "Escape") input.blur();
    return;
  }

  switch (e.code) {
    case "KeyS":
    case "KeyI":
      e.preventDefault();
      input.focus();
      input.select();
      break;

    case "KeyJ":
      scrollDown();
      e.preventDefault();
      break;
    case "KeyK":
      scrollUp();
      e.preventDefault();
      break;
  }
})

inputClearBtn.addEventListener('click', () => {
  input.value = "";
  input.focus();
})

input.addEventListener("keydown", function (e) {
  if (e.key === "Enter") {
    if (e.shiftKey || e.ctrlKey) {
      setSubmiter("text");
      return;
    }
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

const scrollPixel = 300;
function scrollUp() {
  window.scrollBy({
    top: -scrollPixel,
    behavior: 'smooth'
  });
}

function scrollDown() {
  window.scrollBy({
    top: scrollPixel,
    behavior: 'smooth'
  });
}

/*
// TODO: Decided theme is not the hasse for now
// Theme
let currentTheme = "light";
const availabeThemes = ["light", "reading", "dark"];

window.onload = () => {
  document.getElementById(currentTheme).classList.remove('hidden');
}

document.getElementById("themeTouggle").addEventListener('click', () => {
  document.getElementById(currentTheme).classList.add('hidden');
  let next = availabeThemes.indexOf(currentTheme) + 1;
  if (next === availabeThemes.length) {
    next = 0;
  }
  currentTheme = availabeThemes[next];
  document.getElementById(currentTheme).classList.remove('hidden');
})
*/