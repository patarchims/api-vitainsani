# api

# GIT TOKEN
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


upstream stream_backend {
         server tiara.vincentcore.co.id:28444;
}


netstat -tlpn| grep nginx ; ss -tlpn| grep nginx
RUMAH SAKIT HARAPAN



ProxyPreserveHost On
ProxyPass / http://127.0.0.1:8080/
ProxyPassReverse / http://127.0.0.1:8080/
TransferLog /var/log/apache2/yourdomain_access.log
ErrorLog /var/log/apache2/yourdomain_error.log





	query1 := `
	SELECT COALESCE(CONCAT(tgl_periksa,' ',?)+INTERVAL ?*(COUNT(*)) MINUTE,CONCAT(?,' ',?)) AS ELAPSE,
	UNIX_TIMESTAMP(COALESCE(CONCAT(tgl_periksa,' ',?)+INTERVAL ?*(COUNT(*)) MINUTE,CONCAT(?,' ',?)))*1000 AS MILISECOND FROM antrian_ol 
	WHERE kd_dokter=? AND tgl_periksa=? `



LANGKAH-LANGKAH UPLOAD API

masuk ke 
/etc/systemd/system/webapi.service


vi webapi.service
[Unit]
Description=Web Services
ConditionPathExists=/usr/local/bin/api/vincentcore_api-master
After=network.target

[Service]
Type=simple
User=root
Group=root

WorkingDirectory=/usr/local/bin/api/vincentcore_api-master
ExecStart=/usr/local/bin/api/vincentcore_api-master/webapi

Restart=on-failure
RestartSec=10


[Install]
WantedBy=multi-user.target



SETUP NGINX 




SETUP HTTPD

[root@rsvitainsani conf.d]# vi lochost28444-le-ssl.conf
<IfModule mod_ssl.c>
<VirtualHost *:28444>
    ServerAdmin webmaster@rsvitainsani.online
    ServerName rsvitainsani.online
    ServerAlias www.rsvitainsani.online
    DocumentRoot /var/www/html
ErrorLog "logs/ssl28444_error_log"
CustomLog "logs/ssl28444_access_log" combined

Include /etc/letsencrypt/options-ssl-apache.conf
SSLCertificateFile /etc/letsencrypt/live/rsvitainsani.online/cert.pem
SSLCertificateKeyFile /etc/letsencrypt/live/rsvitainsani.online/privkey.pem
SSLCertificateChainFile /etc/letsencrypt/live/rsvitainsani.online/chain.pem
</VirtualHost>
</IfModule>




SETUP NGINX 



[root@192 conf.d]# vi default.conf
server {
  listen 443 ssl;
  listen [::]:443 ssl;

  ssl_certificate /etc/letsencrypt/live/tiara.vincentcore.co.id/fullchain.pem; # managed by Certbot
  ssl_certificate_key /etc/letsencrypt/live/tiara.vincentcore.co.id/privkey.pem; # managed by Certbot
  include /etc/letsencrypt/options-ssl-nginx.conf; # managed by Certbot
  ssl_dhparam /etc/letsencrypt/ssl-dhparams.pem; # managed by Certbot


  location / {
        root /usr/share/nginx/html;
        index index.html index.html;
  }

  error_page  404              /404.html;

  # redirect server error pages to the static page /50x.html
  #
  error_page   500 502 503 504  /50x.html;
  location = /50x.html {
        root   /usr/share/nginx/html;
  }
}

server {
  listen 28444 ssl;
  listen [::]:28444 ssl;

  ssl_certificate /etc/letsencrypt/live/tiara.vincentcore.co.id/fullchain.pem; # managed by Certbot
  ssl_certificate_key /etc/letsencrypt/live/tiara.vincentcore.co.id/privkey.pem; # managed by Certbot
  include /etc/letsencrypt/options-ssl-nginx.conf; # managed by Certbot
  ssl_dhparam /etc/letsencrypt/ssl-dhparams.pem; # managed by Certbot

  #Configures the publicly served root directory
  #Configures the index file to be served
  location / {
        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
        proxy_set_header X-Forwarded-Proto $scheme;
        proxy_set_header X-Real-IP $remote_addr;
        proxy_set_header Host $http_host:28444;
        proxy_pass http://localhost:8080/;
    }
}
