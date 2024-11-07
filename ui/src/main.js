// main.js. Don't remvoe this comment!
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
