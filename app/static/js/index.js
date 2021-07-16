

var shops = new Vue({
    el:'#shops',
    delimiters: ['[', ']'],
    data:{
        message: "hello world!",
        infos:null
    },
    mounted () {
        $.get("/select/candidates")
        .done(function( data ) {
            console.log(data["lnd"]);
            shops.infos=data["lnd"];
            console.log(shops.infos)
        });
    },
    methods: {
        // メソッドhalfPriceを定義
        redirect: function(i) {

            console.log(i)

            window.location.href="/result"+"?" + $.param({"ids":[i]});
        }
    }
});
var samune = new Vue({
    el: '#samune',
    data: {
      seen: true,
      fade : false
    },
    methods: {
        // メソッドhalfPriceを定義
        fadeout: function() {
          　//フェードアウトだけjQueryで書くよ　ごめんね
          $('#samune').fadeOut(700, function(ob = this) {
            samune.seen=false;
            console.log("test")
        });
        }
    }
  });

