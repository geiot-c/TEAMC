
var res = new Vue({
    el:'#res',
    delimiters: ['[(', ')]'],
    data:{
        hoge:"hello",
        params:null,
        selcted:null,
        recommend:null,
        history:null,
        infos:null,
        
    },
    mounted () {
        var query = location.search.substr(1);
        // 「&」で分割して、順に処理する
        var p=[];
        console.log(this.hoge);
        query.split("&").forEach(function (item) {
            // 「=」でパラメーター名と値に分割して、paramsに追加
            var s = item.split("=");
            var k = decodeURIComponent(s[0]);
            var v = decodeURIComponent(s[1]);
            p.push(v);
        });
        console.log(p);
        this.params = p;
        console.log(this.params);
        
        this.selcted=Number(this.params[this.params.length-1]);
        this.recommend=(this.selcted)%5+1;
        if(this.params.length>=2){
            this.history=this.params.slice();
            this.history.pop();
        }
        else{
            this.history=[];
        }
        console.log(this.selcted);
        console.log(this.recommend);
        console.log(this.history);
        $.get("/select/candidates")
        .done(function( data ) {
            console.log(data["lnd"]);
            res.infos=data["lnd"];

        });
    },
    methods: {
        // メソッドhalfPriceを定義
        redirect: function(i) {
            console.log(this.params)
            this.params.push(i)
            window.location.href="/result"+"?" + $.param({"ids":this.params});
        },
        gotop: function() {
            console.log("test");
            window.location.href="/select";
        }
    }
});

