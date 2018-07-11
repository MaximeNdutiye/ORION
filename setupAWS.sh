export AWS_ACCESS_KEY_ID=$(cat accessKeys.csv | awk -F: 'NR==2{print $1}' | awk -F "," '{print $1}')
export AWS_SECRET_ACCESS_KEY=$(cat accessKeys.csv | awk -F: 'NR==2{print $1}' | awk -F "," '{print $2}')
export AWS_DEFAULT_REGION="us-east-1"
echo "Finished setting up keys"