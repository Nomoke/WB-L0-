import 'package:flutter/material.dart';
import 'package:ui/model/order.dart';
import 'package:ui/repository/order.dart';
import 'package:ui/view/widgets/order_card.dart';

class OrderSearchScreen extends StatelessWidget {
  const OrderSearchScreen({super.key});

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        backgroundColor: Theme.of(context).colorScheme.inversePrimary,
        title: const Text('WB(L0) Demo'),
      ),
      body: const Padding(
        padding: EdgeInsets.all(16.0),
        child: _OrderSearchForm(),
      ),
    );
  }
}

class _OrderSearchForm extends StatefulWidget {
  const _OrderSearchForm();

  @override
  _OrderSearchFormState createState() => _OrderSearchFormState();
}

class _OrderSearchFormState extends State<_OrderSearchForm> {
  final TextEditingController _orderIdController = TextEditingController();
  final ord = const OrderRepository();

  @override
  void dispose() {
    _orderIdController.dispose();
    super.dispose();
  }

  @override
  Widget build(BuildContext context) {
    return SingleChildScrollView(
      child: Column(
        children: [
          TextField(
            controller: _orderIdController,
            decoration: const InputDecoration(
              labelText: 'Введите номер заказа',
              border: OutlineInputBorder(),
            ),
          ),
          const SizedBox(height: 16.0),
          ElevatedButton(
            onPressed: () => setState(() {}),
            child: const Text('Поиск'),
          ),
          const SizedBox(height: 30),
          if (_orderIdController.text.isNotEmpty)
            _OrderLoader(
              orderRepository: ord,
              orderId: _orderIdController.text,
            ),
        ],
      ),
    );
  }
}

class _OrderLoader extends StatelessWidget {
  final String orderId;
  final OrderRepository orderRepository;

  const _OrderLoader({
    required this.orderId,
    required this.orderRepository,
  });

  @override
  Widget build(BuildContext context) {
    return FutureBuilder<OrderModel?>(
      future: orderRepository.getOrderById(orderId),
      builder: (context, snapshot) {
        if (snapshot.connectionState == ConnectionState.waiting) {
          return const Center(child: CircularProgressIndicator());
        } else if (snapshot.hasError) {
          return _FetchErrore(snapshot.error);
        } else if (snapshot.hasData) {
          return OrderCardWidget(order: snapshot.data!);
        } else {
          return const Center(child: Text('Но доступно'));
        }
      },
    );
  }
}

class _FetchErrore extends StatelessWidget {
  const _FetchErrore(this.error);

  final Object? error;

  @override
  Widget build(BuildContext context) {
    return Center(
      child: Column(
        children: [
          const Text(
            'Не удалось найти заказ :(',
            style: TextStyle(fontWeight: FontWeight.bold),
          ),
          const SizedBox(
            height: 10,
          ),
          Text('Ошибка:  $error')
        ],
      ),
    );
  }
}
