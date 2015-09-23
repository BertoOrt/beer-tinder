$(document).ready(function(){
    var beerElement = document.getElementById('beer');
    var beer = new Hammer(beerElement);

    function like() {
      $('#beer').addClass('rotate-left').delay(300).fadeOut(1);
      $('.like').show()
      $('.display').css("background-image", "url(http://cliparts.co/cliparts/dc9/rR6/dc9rR6ozi.svg)")
      setTimeout(function () {
        $('.like').hide()
        $('#beer').removeClass('rotate-left').fadeIn(1)
      }, 500)
    }

    function dislike() {
      $('#beer').addClass('rotate-right').delay(300).fadeOut(1);
      $('.dislike').show()
      $('.display').css("background-image", "url(http://www.clker.com/cliparts/1/1/9/2/12065738771352376078Arnoud999_Right_or_wrong_5.svg.med.png)")
      setTimeout(function () {
        $('.dislike').hide()
        $('#beer').removeClass('rotate-right').fadeIn(1)
      }, 500)
    }

    beer.on("panright",function(){
      like()
    });

    beer.on("panleft",function(){
      dislike()
    });

    $(".dislikeButton").click(function(){
     $(this).removeClass("dislikeButton")
     setTimeout(function () {
       $(this).addClass("dislikeButton")
     }.bind(this), 500)
     dislike()
    });

    $('.likeButton').click(function () {
      $(this).removeClass("dislikeButton")
      setTimeout(function () {
        $(this).addClass("likeButton")
      }.bind(this), 500);
      like()
    })

    $('.newBeerForm').submit(function (e) {
      if ($(".verifyName").val() === "" || $(".verifyBrewery").val() === "") {
        e.preventDefault()
      }
    })

    $('.editBeerForm').submit(function (e) {
      if ($(".verifyEditName").val() === "" || $(".verifyEditBrewery").val() === "") {
        e.preventDefault()
      }
    })

    var url = $('.url').val()
    if (url !== "") {
      $(".beerImage").css("background-image", "url(" + url + ")")
    }
});
