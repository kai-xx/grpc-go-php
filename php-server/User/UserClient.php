<?php
// GENERATED CODE -- DO NOT EDIT!

namespace User;

/**
 * 指定 go 的包路径及包名
 * option go_package="app/services;services";
 * 指定 php 的命名空间
 * option php_namespace="App\\Services";
 *
 * User 服务及服务接口的定义
 */
class UserClient extends \Grpc\BaseStub {

    /**
     * @param string $hostname hostname
     * @param array $opts channel options
     * @param \Grpc\Channel $channel (optional) re-use channel object
     */
    public function __construct($hostname, $opts, $channel = null) {
        parent::__construct($hostname, $opts, $channel);
    }

    /**
     * @param \User\UserIndexRequest $argument input argument
     * @param array $metadata metadata
     * @param array $options call options
     * @return \Grpc\UnaryCall
     */
    public function UserIndex(\User\UserIndexRequest $argument,
      $metadata = [], $options = []) {
        return $this->_simpleRequest('/user.User/UserIndex',
        $argument,
        ['\User\UserIndexResponse', 'decode'],
        $metadata, $options);
    }

    /**
     * @param \User\UserViewRequest $argument input argument
     * @param array $metadata metadata
     * @param array $options call options
     * @return \Grpc\UnaryCall
     */
    public function UserView(\User\UserViewRequest $argument,
      $metadata = [], $options = []) {
        return $this->_simpleRequest('/user.User/UserView',
        $argument,
        ['\User\UserViewResponse', 'decode'],
        $metadata, $options);
    }

    /**
     * @param \User\UserPostRequest $argument input argument
     * @param array $metadata metadata
     * @param array $options call options
     * @return \Grpc\UnaryCall
     */
    public function UserPost(\User\UserPostRequest $argument,
      $metadata = [], $options = []) {
        return $this->_simpleRequest('/user.User/UserPost',
        $argument,
        ['\User\UserPostResponse', 'decode'],
        $metadata, $options);
    }

    /**
     * @param \User\UserDeleteRequest $argument input argument
     * @param array $metadata metadata
     * @param array $options call options
     * @return \Grpc\UnaryCall
     */
    public function UserDelete(\User\UserDeleteRequest $argument,
      $metadata = [], $options = []) {
        return $this->_simpleRequest('/user.User/UserDelete',
        $argument,
        ['\User\UserDeleteResponse', 'decode'],
        $metadata, $options);
    }

}
