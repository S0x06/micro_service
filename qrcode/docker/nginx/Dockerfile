# Nginx
#
# VERSION               0.0.1

FROM  nginx:1.12-alpine

MAINTAINER S06 <735594423@qq.com>

EXPOSE 80

volumes:
  - /docker/data/www/service:/var/www/service
  - /docker/data/nginx/conf.d:/etc/nginx/conf.d


CMD ['restart']

