#### AccessWall

###### 인증/암호화 기능을 포함한 서버와 웹 기반의 UI를 중심으로 인증된 사용자들이 서로 암호화 된 파일을 공유하고 제목과 본문이 암호화 된 메일을 주고 받는 것으로 보안에 대한 걱정없이 작업에 대한 효율을 극대화 하는 솔루션입니다.

###### 이 솔루션은 (주)진테크이엔지(jintecheng)의 상용 암호화 솔루션인 AccessWall 3.0 버전중 파일 암호화 와 이메일 암호화 기능을 분리 통합 하여 개인이나 단체가 중요한 정보를 보관 하고 공유 하는데 있어 보안성의 강화 방법을 제공하고자 공개 되고 있습니다.

##### - 설치환경

###### Docker 를 기반으로 Go 언어와 HTML5의 언어로 개발 되었고 MongoDB 를 활용한다. 이는 docker와 Go 언어가 설치되는 어떠한 OS 에서도 서비스가 가능 하다는 것을 의미 하지만 가능하면 LINUX 배포본 중 Server 형의 OS에 설치 운영 할 것을 권장 합니다. 

###### 패키지 예: Centos , Ubuntu LTS Server ..... 등

###### Docker 설치 문서 : <https://docs.docker.com/install/>

###### Golang 설치 문서 : <https://golang.org/doc/install#install>

###### mongodb의 경우는 docker-compose 로 자동 설치

##### - 설치 방법

###### (1) 위에서 설명한 환경의 설정이 완료된 상태에서 설치되기 원하는 디렉토리로 이동 한 뒤 아래 명령을 실행 하여 github.com 으로 부터 소스를 가져 옵니다.

###### $ git clone https://github.com/jintecheng/accesswall

###### (2) 다음 명령으로 설치된 디렉토리로 이동 하여 docker-compose 로 mongodb docker image를 실행 합니다. 

###### $ cd accesswall

###### $ docker-compose up -d 

###### (3) 소스를 build 하여 실행 파일을 생성 합니다. 

###### go build

###### (4) 생성된 실행파일 'accesswall'를 실행 합니다.

###### ./accesswall & 

###### (3), (4) 의 실행 방법은 향후 docker image 화 되어 docker-compose up -d 로 한번에 실행 되게 됩니다. 현재는 개발이 진행 중으로 지금과 같은 수동적인 방법으로 실행 되고 있습니다. 

##### - 솔루션 활용 방법 

###### 향후 업데이트 예정입니다.

##### - 솔루션 총책임자 

###### 최준경 이사

##### - 개발 이슈 관련 연락처

###### <jmryu@jintech2ng.co.kr>

##### - 배포 라이센스 

Copyright Jintecheng Co.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
