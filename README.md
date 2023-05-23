# MuscleApp BackEnd


## Installation

To get started with the project, follow the steps below:

1. Clone the repository:

   ```shell
   git clone https://github.com/kilianp07/MuscleApp.git && cd MuscleApp
   ```

2. Create your .env file and fill it with the required information:
```bash
cp .env.example .env
```

3. Generate JWT Keys:
```bash
echo 'SECRET_KEY="'$(openssl rand -base64 32)'"' >> .env
```
4. Build the project:
```bash
make build
```

6. Run the project:
```bash
./build/MuscleApp
```
