bash env.sh
cd pages && npx tailwindcss -i input.css -o ./dist/output.css
cd ..
cd main && go run *.go 

