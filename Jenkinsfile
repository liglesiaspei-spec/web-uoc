pipeline {
    agent any
    stages {
        stage('Clonar y Verificar Rama') {
            steps {
                echo "Ejecutando pipeline en la rama: ${env.BRANCH_NAME ?: 'main'}"
                sh 'ls -la'
            }
        }
        stage('Construir Imagen Docker') {
            steps {
                script {
                    def tag = (env.BRANCH_NAME == 'dev') ? 'dev' : '1.0'
                    sh "docker build -t luisi333/web-uoc:${tag} ."
                }
            }
        }
        stage('Validar Manifiestos Kubernetes') {
            steps {
                sh 'cat k8s/deployment.yaml && cat k8s/service.yaml'
            }
        }
    }
}
