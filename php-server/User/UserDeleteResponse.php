<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: user.proto

namespace User;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * Generated from protobuf message <code>user.UserDeleteResponse</code>
 */
class UserDeleteResponse extends \Google\Protobuf\Internal\Message
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
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type int $err
     *     @type string $msg
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

}

