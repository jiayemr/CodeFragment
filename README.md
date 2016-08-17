<!doctype html>
<html>
<head>
    <meta charset="utf-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge,chrome=1">
    <meta name="description" content="">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <title>splash screen example</title>
    <link rel="stylesheet" href="http://wayou.github.io/splash-screen-example/nprogress.css">
    <link rel="stylesheet" href="http://wayou.github.io/splash-screen-example/main.css">
    <script src="http://wayou.github.io/splash-screen-example/jquery.min.js"></script>
    <script src="http://wayou.github.io/splash-screen-example/nprogress.js"></script>
    <script type="text/javascript">
        $(function () {
            NProgress.configure({
                template: $('#splash-template').html()
            });
            NProgress.start();

            $('iframe').height($(window).height());
        });
        $(window).load(function () {
            NProgress.done();
        });
    </script>
    <script type="text" id="splash-template">
        <div class="splash card">
            <div role="spinner">
                <div class="spinner-icon"></div>
            </div>
            <p class="lead" style="text-align:center">不要走开，马上回来...</p>
            <div class="progress">
                <div class="mybar" role="bar"></div>
            </div>
        </div>
    </script>
</head>
<body style="overflow: hidden">

<iframe id="iframe" width="100%" src="http://scmship.dev/" frameborder="0"></iframe>
</body>
</html>
