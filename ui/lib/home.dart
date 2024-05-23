import 'package:flutter/material.dart';
import 'package:ui/view/search_screen.dart';

class WBOrdersApp extends StatelessWidget {
  const WBOrdersApp({super.key});

  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      debugShowCheckedModeBanner: false,
      title: 'WB(L0) Demo',
      theme: ThemeData(
        colorScheme: ColorScheme.fromSeed(seedColor: Colors.deepPurple),
        useMaterial3: true,
      ),
      home: const OrderSearchScreen(),
    );
  }
}
