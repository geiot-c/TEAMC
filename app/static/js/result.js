
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
        self_info:"test",

        
    },
    mounted () {
        var query = location.search.substr(1);
        // 「&」で分割して、順に処理する
        var p=[];
        /*query.split("&").forEach(function (item) {
            // 「=」でパラメーター名と値に分割して、paramsに追加
            var s = item.split("=");
            var k = decodeURIComponent(s[0]);
            var v = decodeURIComponent(s[1]);
            p.push(v);
        });*/
        this.selcted=Number(location.pathname.split('/')[2]);
        console.log(location.href+"/status");
        $.get(location.href+"/status")
        .done(function( data ) {
            console.log(data);
            res.infos=data;
            res.set_map(data)
        });
        this.recommend=(this.selcted)%13+1;
        
        console.log(this.selcted);
        console.log(this.infos["recommend"]);
        console.log(this.history);

        $.get("/select/candidates")
        .done(function( data ) {
            console.log(data["lnd"]);
            res.infos=data["lnd"];

        });
    },
    methods: {
        
        redirect: function(i) {
            window.location.href="/result/"+i;
        },
        gotop: function() {
            console.log("test");
            window.location.href="/select";
        },
        set_map: function(i) {
            var map = L.map('mapid', {
            center: [i["latitude"], i["longitude"]],
            zoom: 12,
             }); 
        // OpenStreetMap から地図画像を読み込む
        
        var tileLayer = L.tileLayer('https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png', {
            attribution: '© <a href="http://osm.org/copyright">OpenStreetMap</a> contributors, <a href="http://creativecommons.org/licenses/by-sa/2.0/">CC-BY-SA</a>',
        });
        tileLayer.addTo(map);
        
        // マーカー画像の場所を指定する
        var icon_self_pos = L.icon({
            iconUrl: '/image/self_pos.png',
            iconSize:     [32, 32], // アイコンのサイズ
            iconAnchor:   [16, 52], // マーカーの位置に対応するアイコンの位置
            popupAnchor:  [0, -55] // ポップアップを開く基準
        });

        var icon_recommend_pos = L.icon({
            iconUrl: '/image/recommend_pos.png',
            iconSize:     [32, 32], // アイコンのサイズ
            iconAnchor:   [16, 52], // マーカーの位置に対応するアイコンの位置
            popupAnchor:  [0, -55] // ポップアップを開く基準
        });
        L.marker([i["latitude"], i["longitude"]], {icon: icon_self_pos}).addTo(map)
            .bindPopup(`${i["name"]}<br><img  src="/image/${i["id"]}.jpg" height="100">`)
            .openPopup(); 
                    
        $.each(i["recommend"], function(index, value){
            let popup = `${value["name"]}<br>
                <button class ="logo" onclick="location.href='/result/${value["id"]}'"> 
                <img src="/image/${value["id"]}.jpg" height="100">
                </button>
                <br>
                おすすめポイント<br>
                ${value["intro"]}
                `
            L.marker([value["latitude"], value["longitude"]],{icon: icon_recommend_pos}).addTo(map)
                .bindPopup(popup);
            });
        }
    }  
});
