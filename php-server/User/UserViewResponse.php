<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: user.proto

namespace User;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * Generated from protobuf message <code>user.UserViewResponse</code>
 */
class UserViewResponse extends \Google\Protobuf\Internal\Message
{
    /**
     * Generated from protobuf field <code>int32 err = 1;</code>
     */
    protected $err = 0;
    /**
     * Generated from protobuf field <code>string msg = 2;</code>
     */
    protected $msg = '';
    /**
     * 返回一个 UserEntity 对象
     *
     * Generated from protobuf field <code>.user.UserEntity data = 3;</code>
     */
    protected $data = null;

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type int $err
     *     @type string $msg
     *     @type \User\UserEntity $data
     *           返回一个 UserEntity 对象
     * }
     */
    public function __construct($data = NULL) {
        \GPBMetadata\User::initOnce();
        parent::__construct($data);
    }

    /**
     * Generated from protobuf field <code>int32 err = 1;</code>
     * @return int
     */
    public function getErr()
    {
        return $this->err;
    }

    /**
     * Generated from protobuf field <code>int32 err = 1;</code>
     * @param int $var
     * @return $this
     */
    public function setErr($var)
    {
        GPBUtil::checkInt32($var);
        $this->err = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>string msg = 2;</code>
     * @return string
     */
    public function getMsg()
    {
        return $this->msg;
    }

    /**
     * Generated from protobuf field <code>string msg = 2;</code>
     * @param string $var
     * @return $this
     */
    public function setMsg($var)
    {
        GPBUtil::checkString($var, True);
        $this->msg = $var;

        return $this;
    }

    /**
     * 返回一个 UserEntity 对象
     *
     * Generated from protobuf field <code>.user.UserEntity data = 3;</code>
     * @return \User\UserEntity|null
     */
    public function getData()
    {
        return $this->data;
    }

    public function hasData()
    {
        return isset($this->data);
    }

    public function clearData()
    {
        unset($this->data);
    }

    /**
     * 返回一个 UserEntity 对象
     *
     * Generated from protobuf field <code>.user.UserEntity data = 3;</code>
     * @param \User\UserEntity $var
     * @return $this
     */
    public function setData($var)
    {
        GPBUtil::checkMessage($var, \User\UserEntity::class);
        $this->data = $var;

        return $this;
    }

}
