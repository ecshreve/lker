<!doctype html>
<html lang="en">

  <head>
    <meta charset="utf-8">
    <meta name="viewport" content="width=device-width, initial-scale=1, shrink-to-fit=no">
    <meta name="description" content="">
    <script>
      var targ = {{ .B }};
      var start = new Date().getTime();
      var x = setInterval(function() {
        var now = new Date().getTime();
        diff = now - start
        targ = targ - diff
        document.getElementById("count").innerHTML = "<strong>" + targ.toString() + "<strong>";
        if (targ < 100) {
            clearInterval(x);
            document.getElementById("count").innerHTML = "";
        }
        start = now
      }, 1000);
    </script>
    <title>guess</title>

    <link rel="canonical" href="https://getbootstrap.com/docs/4.6/examples/starter-template/">



    <!-- Bootstrap core CSS -->
    <link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/bootstrap@4.6.2/dist/css/bootstrap.min.css" integrity="sha384-xOolHFLEh07PJGoPkLv1IbcEPTNtaed2xpHsD9ESMhqIYd0nLMwNLD69Npy4HI+N" crossorigin="anonymous">


    <style>
      .bd-placeholder-img {
        font-size: 1.125rem;
        text-anchor: middle;
        -webkit-user-select: none;
        -moz-user-select: none;
        -ms-user-select: none;
        user-select: none;
      }

      @media (min-width: 768px) {
        .bd-placeholder-img-lg {
          font-size: 3.5rem;
        }
      }
    </style>


    <!-- Custom styles for this template -->
    <link href="/static/style.css" rel="stylesheet">

    <!-- Matomo -->
    <script>
      var _paq = window._paq = window._paq || [];
      /* tracker methods like "setCustomDimension" should be called before "trackPageView" */
      _paq.push(['trackPageView']);
      _paq.push(['enableLinkTracking']);
      (function() {
        var u="//172.18.0.3/";
        _paq.push(['setTrackerUrl', u+'matomo.php']);
        _paq.push(['setSiteId', '1']);
        var d=document, g=d.createElement('script'), s=d.getElementsByTagName('script')[0];
        g.async=true; g.src=u+'matomo.js'; s.parentNode.insertBefore(g,s);
      })();
    </script>
    <!-- End Matomo Code -->

  </head>

  <body>

    <main
      role="main"
      class="container"
    >
      <div class="wave"></div>
      <div id="count"></div>
      <div class="wave"></div>

      <div
        class="row"
        style="margin-top: 25px;"
      >
        <div class="col-2"></div>
        <div class="col-4">
          <form>
            <div>
              <input
                type="text"
                name="guessbox"
                id="guessbox"
              >
              <button type="submit" class="btn btn-primary">guess</button>
            </div>
          </form>
        </div>
        <div class="col-6">
          <div id="scroll-container">
            <div id="scroll-text">
              {{range .G}}{{ . }}</br>{{end}}
            </div>
          </div>
        </div>
      </div>


      <div class="row">
        <div class="col-2"></div>
        <div class="col-8">
          <ul class="cloud">
            {{ range .W }}
            <li><a
                href="#"
                data-weight="{{ .Rank }}"
              >{{ .Val }}</a></li>{{end}}
          </ul>
        </div>
        <div class="col-2"></div>
      </div>





    </main><!-- /.container -->


    <script src="https://cdn.jsdelivr.net/npm/jquery@3.5.1/dist/jquery.slim.min.js" integrity="sha384-DfXdz2htPH0lsSSs5nCTpuj/zy4C+OGpamoFVy38MVBnE+IbbVYUew+OrCXaRkfj" crossorigin="anonymous"></script>
    <script>window.jQuery || document.write('<script src="../assets/js/vendor/jquery.slim.min.js"><\/script>')</script>
    <script src="../assets/dist/js/bootstrap.bundle.min.js"></script>


  </body>

</html>