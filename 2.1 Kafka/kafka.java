public class ProducerRecord<K, V> {
    // 主题
    private final String topic;
    // 分区号
    private final Integer partition;
    // 消息头部
    private final Headers headers;
    // 键
    // 非必输 可以用来计算分区号进而可以让消息发往特定的分区
    private final K key;
    // 消息体 一般不为空
    private final V value;
    // 消息的时间戳
    private final Long timestamp;
}