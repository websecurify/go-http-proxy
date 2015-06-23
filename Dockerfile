FROM scratch

# ---
# ---
# ---

COPY go-http-proxy /

# ---
# ---
# ---

EXPOSE 8080

# ---
# ---
# ---

ENTRYPOINT ["/go-http-proxy"]

# ---
