let up = document.getElementById('up-btn');

window.addEventListener('scroll', function() {
    up.hidden = (window.pageYOffset < document.documentElement.clientHeight)
  });
