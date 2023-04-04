function changeTab(film_type, btn_name) {
  document.getElementById(film_type).style.display = "";

  if(film_type == "film") {
    document.getElementById('iphone').style.display = "none";
  } else {
    document.getElementById('film').style.display = "none";
  }
}

document.getElementById('iphone').style.display = "none";

window.onbeforeunload = function () {
  window.scrollTo(0, 0);
}

const loader = document.querySelector(".loader");
window.onload = function(){
  setTimeout(function(){
      loader.style.opacity = "0";
      setTimeout(function(){
          loader.style.display = "none";
          }, 500);
      },1500);
}
