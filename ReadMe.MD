docker build --platform linux/amd64 -t product-inventory .
docker tag product-inventory:latest gcr.io/essential-rig-435514-k4/product-inventory:latest
docker push gcr.io/essential-rig-435514-k4/product-inventory:latest

helm list
helm uninstall product-inv-chart
helm install product-inv-chart /Users/saifalisk/Desktop/capgemini/helm/product-inv-chart

Install Postgres using our own values.yaml
cd postgres
helm install my-postgresql bitnami/postgresql -f values.yaml --namespace default