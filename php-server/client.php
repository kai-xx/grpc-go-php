<?php
/*
 *
 * Copyright 2015 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

// php:generate protoc --proto_path=./../protos   --php_out=./   --grpc_out=./ --plugin=protoc-gen-grpc=./../../bins/opt/grpc_php_plugin ./../protos/helloworld.proto


require dirname(__FILE__).'/vendor/autoload.php';

$hostname = !empty($argv[2]) ? $argv[2] : 'localhost:50051';

$client = new \User\UserClient($hostname, [
    'credentials' => Grpc\ChannelCredentials::createInsecure(),
]);
$indexRequest = new \User\UserIndexRequest();
$indexRequest->setPage(1);
$indexRequest->setPageSize(12);
list($response, $status) = $client->UserIndex($indexRequest)->wait();
if ($status->code !== Grpc\STATUS_OK) {
    echo "ERROR: " . $status->code . ", " . $status->details . PHP_EOL;
}
/**
 * @var $value \User\UserEntity
 */
foreach ($response->getData() as $key=>$value) {
    echo "name =>{$value->getName()} --- age =>{$value->getAge()} " . PHP_EOL;
}
echo "index {$response->getMsg()}" . PHP_EOL;
$viewRequest = new \User\UserViewRequest();
$viewRequest->setUid(1);

list($response, $status) = $client->UserView($viewRequest)->wait();
if ($status->code !== Grpc\STATUS_OK) {
    echo "ERROR: " . $status->code . ", " . $status->details . PHP_EOL;
}
$view = $response->getData();
echo "name =>{$view->getName()} --- age =>{$view->getAge()} " . PHP_EOL;
echo "view {$response->getMsg()}" . PHP_EOL;
$postRequest = new \User\UserPostRequest();
$postRequest->setAge(12);
$postRequest->setName("12312");
$postRequest->setPassword("324523");
list($response, $status) = $client->UserPost($postRequest)->wait();
if ($status->code !== Grpc\STATUS_OK) {
    echo "ERROR: " . $status->code . ", " . $status->details . PHP_EOL;
}
echo "post {$response->getMsg()}" . PHP_EOL;

