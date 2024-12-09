// main.js. Don't remvoe this comment!
const inputClearBtn = document.getElementById("inputClear")
const input = document.getElementById("input");
const inputSugg = document.getElementById("inputSuggestions");

function isSuggHidden() { return inputSugg.classList.contains('hidden'); }

window.addEventListener('scroll', function (e) {
  if (!isSuggHidden()) setTimeout(() => inputSugg.classList.add('hidden'), 100);
})


document.addEventListener('click', (e) => {
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
    if (input.value !== "") {
      const url = `/sugg?w=${encodeURIComponent(input.value)}`;
      console.log("req", url)
      const res = await fetch(url, { method: "POST" });
      if (res.ok) {
        const data = await res.text();
        if (data) {
          inputSugg.innerHTML = data;
          inputSugg.classList.remove("hidden");
          return;
        }
      }
    }
    // the func would return if no error was encountered. the server should error incase no data
    inputSugg.classList.add("hidden");
    inputSugg.innerHTML = "";

  }, 400);
})

// ----------- keyboard shoutcuts ----------
document.addEventListener('keydown', (e) => {
  if (document.activeElement === input) {
    if (e.code === "Escape") input.blur();
    return;
  }


  // no composite key
  if (/*e.shiftKey ||*/ e.ctrlKey) {
    return;
  }

  switch (e.code) {
    case "KeyS":
      e.preventDefault();
      input.focus();
      input.setSelectionRange(input.value.length, input.value.length);
      break;
    case "KeyI":
      e.preventDefault();
      input.focus();
      input.select();
      break;

    case "KeyT":
      e.preventDefault();
      if (window.location.search) {
        loadPage(`/t${window.location.search}`);
      }
      break;
    case "KeyR":
      e.preventDefault();
      if (window.location.search) {
        loadPage(`/r${window.location.search}`);
      }
      break;

    case "KeyF":
      e.preventDefault();
      loadPage("/roots");
      break;
    case "KeyH":
      e.preventDefault();
      loadPage("/");
      break;

    case "KeyJ":
      e.preventDefault();
      scrollDown();
      break;
    case "KeyK":
      e.preventDefault();
      scrollUp();
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