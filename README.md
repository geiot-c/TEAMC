# TEAMC

## Docker操作
起動(goサーバーも自動起動)：  
docker-compose up -d --build  
コンテナ内確認(bashログイン):   
docker exec -it app bash  
削除：  
docker-compose down -v  
ログ確認:  
docker-compose logs  
コンテナ生死確認  
docker-compose ps  

## コンテナ内
bashログアウト:  
$exit

## 動作確認
http://34.127.0.137:8080/(任意)
