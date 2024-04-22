FROM debian:stable-slim

COPY nutrition-calculator /bin/nutrition-calculator

CMD ["/bin/nutrition-calculator"]