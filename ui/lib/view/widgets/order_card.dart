import 'package:flutter/material.dart';
import 'package:ui/model/order.dart';

class OrderCardWidget extends StatelessWidget {
  const OrderCardWidget({super.key, required this.order});

  final OrderModel order;

  @override
  Widget build(BuildContext context) {
    return Card(
      elevation: 4,
      margin: const EdgeInsets.all(8),
      child: Padding(
        padding: const EdgeInsets.all(12),
        child: Column(
          crossAxisAlignment: CrossAxisAlignment.start,
          children: [
            Text(
              'Номер заказа: ${order.id}',
              style: const TextStyle(
                fontWeight: FontWeight.bold,
                fontSize: 18,
              ),
            ),
            const SizedBox(height: 8),
            Text('Дата создания: ${order.dateCreated}'),
            Text('Номер отслеживания: ${order.trackNumber}'),
            Text('Служба доставки: ${order.deliveryService}'),
            Text('ID доставки: ${order.deliveryId}'),
            Text('ID платежа: ${order.paymentId}'),
            Text('Сумма заказа: ${order.payment.amount} RUB'),
            Text('Стоимость доставки: ${order.payment.deliveryCost} RUB'),
            Text('Итоговая стоимость: ${order.payment.goodsTotal} RUB'),
            const SizedBox(height: 16),
            const Text(
              'Товары:',
              style: TextStyle(
                fontWeight: FontWeight.bold,
                fontSize: 16,
              ),
            ),
            const SizedBox(height: 8),
            SizedBox(
              height: 120,
              child: ListView.builder(
                scrollDirection: Axis.horizontal,
                itemCount: order.items.length,
                itemBuilder: (context, index) {
                  final item = order.items[index];
                  return Container(
                    margin: const EdgeInsets.only(right: 12),
                    padding: const EdgeInsets.all(8),
                    decoration: BoxDecoration(
                      border: Border.all(
                        color: Colors.grey.shade300,
                      ),
                      borderRadius: BorderRadius.circular(8),
                    ),
                    child: Column(
                      crossAxisAlignment: CrossAxisAlignment.start,
                      children: [
                        Text(
                          'Название: ${item.name}',
                          style: const TextStyle(
                            fontWeight: FontWeight.bold,
                          ),
                        ),
                        Text('Цена: ${item.price} RUB'),
                        Text('Размер: ${item.size}'),
                        Text('Бренд: ${item.brand}'),
                        Text('Статус: ${item.status}'),
                      ],
                    ),
                  );
                },
              ),
            )
          ],
        ),
      ),
    );
  }
}
