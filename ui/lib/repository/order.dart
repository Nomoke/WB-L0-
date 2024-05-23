import 'dart:convert';

import 'package:http/http.dart' as http;
import 'package:ui/model/order.dart';

class OrderRepository {
  const OrderRepository();

  final String baseUrl = 'http://localhost:8080';

  Future<OrderModel?> getOrderById(String id) async {
    var url = Uri.parse('$baseUrl/$id');

    await Future.delayed(const Duration(milliseconds: 360));

    try {
      var response = await http.get(url);
      if (response.statusCode == 200) {
        var jsonResponse = json.decode(response.body);
        return OrderModel.fromJson(jsonResponse);
      } else {
        throw Exception(
            'Failed to load order. Status code: ${response.statusCode}');
      }
    } catch (e) {
      throw Exception('Failed to load order: $e');
    }
  }
}
