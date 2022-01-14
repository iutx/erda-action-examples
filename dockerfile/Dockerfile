FROM registry.cn-shanghai.aliyuncs.com/viper/rust-musl-builder:1.56.0 as rust-musl-builder

USER root

WORKDIR /build

COPY . /build


RUN mkdir -p /root/.cargo && \
    echo '[source.crates-io]\nregistry = "https://github.com/rust-lang/crates.io-index"\nreplace-with = "ustc"\n[source.ustc]\nregistry = "git://mirrors.ustc.edu.cn/crates.io-index"' >> /root/.cargo/config && \
    cargo build --release

FROM neurodebian:stretch
LABEL maintainer="iutx<root@viper.run>"
ENV LANG=en_US.UTF-8

WORKDIR /app

COPY --from=rust-musl-builder /build/target/release/axum-demo /app/axum-demo

EXPOSE 3000

CMD ["./axum-demo"]