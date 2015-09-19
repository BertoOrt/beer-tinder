$(document).ready(function(){

    function like() {
      $('.beer').addClass('rotate-left').delay(700).fadeOut(1);
      $('.like').show()
      $('.display').css("background-image", "url(http://www.clker.com/cliparts/2/k/n/l/C/Q/transparent-green-checkmark-hi.png)")
      setTimeout(function () {
        $('.like').hide()
        $('.beer').removeClass('rotate-left').fadeIn(1)
      }, 1500)
    }

    function dislike() {
      $('.beer').addClass('rotate-right').delay(700).fadeOut(1);
      $('.dislike').show()
      $('.display').css("background-image", "url(http://www.clipartbest.com/cliparts/yTo/xdd/yToxddATE.png)")
      setTimeout(function () {
        $('.dislike').hide()
        $('.beer').removeClass('rotate-right').fadeIn(1)
      }, 1500)
    }

    $(".beer").on("swiperight",function(){
      like()
    });

    $(".beer").on("swipeleft",function(){
      dislike()
    });

    $(".dislikeButton").click(function(){
     $(this).removeClass("dislikeButton")
     setTimeout(function () {
       $(this).addClass("dislikeButton")
     }.bind(this), 1500)
     dislike()
    });

    $('.likeButton').click(function () {
      setTimeout(function () {
        $(this).addClass("likeButton")
      }.bind(this), 1500);
      like()
    })

});
