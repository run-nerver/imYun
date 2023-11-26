# node镜像
FROM node:16.14.0 as build-stage

RUN echo "-------------------- web环境配置 --------------------"

# 指定接下来的工作路径为/app  - 类似于cd命令
WORKDIR /app
# 拷贝前端项目到app目录下
ADD . /app/
# RUN chmod +x docker_init.sh
# RUN sh docker_init.sh
# 设置淘宝npm镜像
RUN npm install -g cnpm --registry=http://registry.npmmirror.com
# 安装依赖
RUN cnpm install
RUN cnpm install --save body-parser
RUN cnpm install --save core-js
RUN cnpm install --save svg-baker-runtime
# 打包 - 目的：丢到nginx下跑
RUN cnpm run build:prod

# 前端项目运行命令
#CMD ["npm","run","start"]


# ======================== 上：npm打包  下：nginx运行 ========================
# nginx镜像
FROM nginx:1.15.3-alpine as production-stage

# 移除nginx容器的default.conf文件、nginx配置文件
RUN rm /etc/nginx/conf.d/default.conf
RUN rm /etc/nginx/nginx.conf
# 把主机的nginx.conf文件复制到nginx容器的/etc/nginx文件夹下
COPY ./nginx.conf /etc/nginx/
# 拷贝前端vue项目打包后生成的文件到nginx下运行
COPY --from=build-stage /app/dist /usr/share/nginx/html
# WORKDIR /app
EXPOSE 9527

# 注：CMD不同于RUN，CMD用于指定在容器启动时所要执行的命令，而RUN用于指定镜像构建时所要执行的命令。
#    RUN指令创建的中间镜像会被缓存，并会在下次构建中使用。如果不想使用这些缓存镜像，可以在构建时指定--no-cache参数，如：docker build --no-cache

# 使用daemon off的方式将nginx运行在前台保证镜像不至于退出
CMD ["nginx", "-g", "daemon off;"]
# CMD /app/docker_init.sh
# CMD ["/app/docker_init.sh"]
