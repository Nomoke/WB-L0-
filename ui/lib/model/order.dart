class OrderModel {
  final String id;
  final DateTime dateCreated;
  final String trackNumber;
  final String deliveryService;
  final int deliveryId;
  final int paymentId;
  final String locale;
  final String internalSignature;
  final String customerId;
  final String shardKey;
  final int smId;
  final String oofShard;
  final Delivery delivery;
  final Payment payment;
  final List<OrderItem> items;

  OrderModel({
    required this.id,
    required this.dateCreated,
    required this.trackNumber,
    required this.deliveryService,
    required this.deliveryId,
    required this.paymentId,
    required this.locale,
    required this.internalSignature,
    required this.customerId,
    required this.shardKey,
    required this.smId,
    required this.oofShard,
    required this.delivery,
    required this.payment,
    required this.items,
  });

  factory OrderModel.fromJson(Map<String, dynamic> json) {
    return OrderModel(
      id: json['order_uid'],
      dateCreated: DateTime.parse(json['date_created']),
      trackNumber: json['track_number'],
      deliveryService: json['delivery_service'],
      deliveryId: json['delivery']['id'],
      paymentId: json['payment']['id'],
      locale: json['locale'],
      internalSignature: json['internal_signature'],
      customerId: json['customer_id'],
      shardKey: json['shard_key'],
      smId: json['sm_id'],
      oofShard: json['oof_shard'],
      delivery: Delivery.fromJson(json['delivery']),
      payment: Payment.fromJson(json['payment']),
      items: List<OrderItem>.from(
          json['items'].map((item) => OrderItem.fromJson(item))),
    );
  }
}

class Delivery {
  final int id;
  final String name;
  final String phone;
  final String zip;
  final String city;
  final String address;
  final String region;
  final String email;

  Delivery({
    required this.id,
    required this.name,
    required this.phone,
    required this.zip,
    required this.city,
    required this.address,
    required this.region,
    required this.email,
  });

  factory Delivery.fromJson(Map<String, dynamic> json) {
    return Delivery(
      id: json['id'],
      name: json['name'],
      phone: json['phone'],
      zip: json['zip'],
      city: json['city'],
      address: json['address'],
      region: json['region'],
      email: json['email'],
    );
  }
}

class Payment {
  final int id;
  final String transaction;
  final String requestId;
  final String currency;
  final String provider;
  final int amount;
  final DateTime paymentDt;
  final String bank;
  final int deliveryCost;
  final int goodsTotal;
  final int customFee;

  Payment({
    required this.id,
    required this.transaction,
    required this.requestId,
    required this.currency,
    required this.provider,
    required this.amount,
    required this.paymentDt,
    required this.bank,
    required this.deliveryCost,
    required this.goodsTotal,
    required this.customFee,
  });

  factory Payment.fromJson(Map<String, dynamic> json) {
    return Payment(
      id: json['id'],
      transaction: json['transaction'],
      requestId: json['request_id'],
      currency: json['currency'],
      provider: json['provider'],
      amount: json['amount'],
      paymentDt: DateTime.parse(json['payment_dt']),
      bank: json['bank'],
      deliveryCost: json['delivery_cost'],
      goodsTotal: json['goods_total'],
      customFee: json['custom_fee'],
    );
  }
}

class OrderItem {
  final int chrtId;
  final String trackNumber;
  final int price;
  final String rid;
  final String name;
  final int sale;
  final String size;
  final int totalPrice;
  final int nmId;
  final String brand;
  final int status;

  OrderItem({
    required this.chrtId,
    required this.trackNumber,
    required this.price,
    required this.rid,
    required this.name,
    required this.sale,
    required this.size,
    required this.totalPrice,
    required this.nmId,
    required this.brand,
    required this.status,
  });

  factory OrderItem.fromJson(Map<String, dynamic> json) {
    return OrderItem(
      chrtId: json['chrt_id'],
      trackNumber: json['track_number'],
      price: json['price'],
      rid: json['rid'],
      name: json['name'],
      sale: json['sale'],
      size: json['size'],
      totalPrice: json['total_price'],
      nmId: json['nm_id'],
      brand: json['brand'],
      status: json['status'],
    );
  }
}
