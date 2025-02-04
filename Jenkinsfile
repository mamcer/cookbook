pipeline {
	agent any
	environment {
		PATH = '${env.PATH}'
	} 
	stages {
		stage("build") {
			steps {
				sh '/usr/local/go/bin/go version'
			}
		}
	}
    post {
        success {
            emailext (
                to: 'mario.moreno@live.com',
                subject: "SUCCESS: Job '${env.JOB_NAME} [${env.BUILD_NUMBER}]'",
                body: """<p>Job '${env.JOB_NAME} [${env.BUILD_NUMBER}]' succeeded.</p><p>Check console output at <a href='${env.BUILD_URL}'>${env.BUILD_URL}</a></p>""",
                mimeType: 'text/html'
            )
        }
        failure {
            emailext (
                to: 'mario.moreno@live.com',
                subject: "FAILURE: Job '${env.JOB_NAME} [${env.BUILD_NUMBER}]'",
                body: """<p>Job '${env.JOB_NAME} [${env.BUILD_NUMBER}]' failed.</p><p>Check console output at <a href='${env.BUILD_URL}'>${env.BUILD_URL}</a></p>""",
                mimeType: 'text/html'
            )
        }
        always {
            cleanWs()
        }
    }
}
