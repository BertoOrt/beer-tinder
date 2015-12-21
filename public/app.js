$(document).ready(function(){

  // **************
  // HammerJS Setup
  // **************

  var beerElement = document.getElementById('beer');
  var beer = new Hammer(beerElement);
  beer.get('pan').set({ direction: Hammer.DIRECTION_ALL, threshold: 1 });

  // ***************
  // Swipe Listeners
  // ***************

  function like() {
    $('#beer').addClass('rotate-left').delay(300).fadeOut(1);
    $('.like').show()
    $('.display').css("background-image", "url(http://cliparts.co/cliparts/dc9/rR6/dc9rR6ozi.svg)")
    $.post("/tally/" + $(".currentTimestamp").val(), {"score": "Up"})
      .done(function (response) {
      });
    setTimeout(function () {
      $('.like').hide()
      $('#beer').removeClass('rotate-left').fadeIn(1)
    }, 500)
  }

  function dislike() {
    $('#beer').addClass('rotate-right').delay(300).fadeOut(1);
    $('.dislike').show()
    $('.display').css("background-image", "url(http://www.clker.com/cliparts/1/1/9/2/12065738771352376078Arnoud999_Right_or_wrong_5.svg.med.png)")
    $.post("/tally/" + $(".currentTimestamp").val(), {"score": "Down"})
      .done(function (response) {
      });
    setTimeout(function () {
      $('.dislike').hide()
      $('#beer').removeClass('rotate-right').fadeIn(1)
    }, 500)
  }

  beer.on("swiperight", function(){
    like()
  });

  beer.on("swipeleft", function(){
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

  //**********
  //*Get Info*
  //**********

  $('.beerName').on('click', function (e) {
    $.get('/current')
      .done(function (result) {
        $('.infoUp span').html(result.data.Up);
        $('.infoDown span').html(result.data.Down);
      })
  })

  //**********
  //Validation
  //**********

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

  //***********
  // Image URL
  //***********

  var url = $('.url').val()
  if (url !== "") {
    $(".beer").css("background-image", "url(" + url + ")")
  }
});
