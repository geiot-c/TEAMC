

//写真clickでリンクを移動
var shops = new Vue({
    el:'#shops',
    //{{}}を[()]に変更
    delimiters: ['[(', ')]'],
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
        redirect: function(i) {
            console.log(i)
            //window.location.href="/result"+"?" + $.param({"ids":[i]});
            window.location.href="/result"+"/" + i;
        }
    }
});
//最初のフェードアウトの設定
var samune = new Vue({
    el: '#samune',
    data: {
      seen: true,
      fade : false
    },
    methods: {
        fadeout: function() {
          　//フェードアウトだけjQueryで書くよ　ごめんね
          $('#samune').fadeOut(700, function(ob = this) {
            samune.seen=false;
            console.log("test")
        });
        }
    }
  });

