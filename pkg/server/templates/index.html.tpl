<!doctype html>
<html lang="en">

  <head>
    <meta charset="utf-8">
    <meta name="viewport" content="width=device-width, initial-scale=1, shrink-to-fit=no">
    <meta name="description" content="guess">

    <!-- countdown from initial value -->
    <script>
      var targ = {{ .B }};
      var x = setInterval(function() {
        targ = targ - (1011)
        document.getElementById("count").innerHTML = "<strong>" + targ.toString() + "<strong>";
      }, 1011);
    </script>

    <title>guess</title>
    <!-- Bootstrap core CSS -->
    <link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/bootstrap@4.6.2/dist/css/bootstrap.min.css" integrity="sha384-xOolHFLEh07PJGoPkLv1IbcEPTNtaed2xpHsD9ESMhqIYd0nLMwNLD69Npy4HI+N" crossorigin="anonymous">

    <!-- Custom styles -->
    <link href="/style.css" rel="stylesheet">
  </head>

  <body>
    <main class="container">

      <div class="wave"></div>
      <div class="wave"></div>

      <!-- countdown value -->
      <div id="count"></div>

      <div
        class="row"
        style="margin-top: 25px;"
      >
        <div class="col-2"></div>

        <!-- guess form -->
        <div class="col-4">
          <form
            method="POST"
            action="/"
          >
            <div>
              <input
                type="text"
                name="guessbox"
                id="guessbox"
              >
              <button type="submit" {{ .D }} class="btn-custom">guess</button>
            </div>
          </form>
        </div>

        <!-- scrolling text -->
        <div class="col-6">
          <div id="scroll-container">
            <div id="scroll-text">
              {{range .G}}{{ . }}</br>{{end}}
            </div>
          </div>
        </div>
      </div>

      <!-- letter cloud -->
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
    </main>
  </body>

</html>