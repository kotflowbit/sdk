<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: rpc/block.proto

namespace KotFlowBlock;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * Generated from protobuf message <code>KotFlowBlock.ReplyLatest</code>
 */
class ReplyLatest extends \Google\Protobuf\Internal\Message
{
    /**
     * Generated from protobuf field <code>bytes Height = 1;</code>
     */
    protected $Height = '';
    /**
     * Generated from protobuf field <code>bytes Hash = 2;</code>
     */
    protected $Hash = '';

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type string $Height
     *     @type string $Hash
     * }
     */
    public function __construct($data = NULL) {
        \GPBMetadata\Rpc\Block::initOnce();
        parent::__construct($data);
    }

    /**
     * Generated from protobuf field <code>bytes Height = 1;</code>
     * @return string
     */
    public function getHeight()
    {
        return $this->Height;
    }

    /**
     * Generated from protobuf field <code>bytes Height = 1;</code>
     * @param string $var
     * @return $this
     */
    public function setHeight($var)
    {
        GPBUtil::checkString($var, False);
        $this->Height = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>bytes Hash = 2;</code>
     * @return string
     */
    public function getHash()
    {
        return $this->Hash;
    }

    /**
     * Generated from protobuf field <code>bytes Hash = 2;</code>
     * @param string $var
     * @return $this
     */
    public function setHash($var)
    {
        GPBUtil::checkString($var, False);
        $this->Hash = $var;

        return $this;
    }

}

