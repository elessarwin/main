import 'dart:async';

export './bloc/app_config.dart';
export './bloc/data.dart';
export './config/chat_config.dart';
export './services/auth_user_service.dart';
export './services/cache_service.dart';
export './services/item_repository.dart';
export './services/network.dart';
export './services/storage.dart';
export './services/user.dart';
export './utils/apptheme.dart';
export './utils/utils.dart';
export './widgets/avatar_loader.dart';
export './widgets/home_screen.dart';
export './widgets/login_screen.dart';
export './widgets/settings.dart';
export './widgets/splash_screen.dart';
export './widgets/title_widget.dart';
export './widgets/video_player.dart';
export './widgets/web_screen/bottom_nav.dart';
export './widgets/web_screen/layouts.dart';
export './widgets/web_screen/simple_route.dart';

Completer<dynamic> storageServiceReadyCompleter = Completer<dynamic>();
Completer<dynamic> networkServiceReadyCompleter = Completer<dynamic>();
Completer<dynamic> userServiceReadyCompleter = Completer<dynamic>();
Completer<dynamic> authUserServiceReadyCompleter = Completer<dynamic>();
Completer<dynamic> cacheServiceReadyCompleter = Completer<dynamic>();
Future<dynamic> get storageReady => storageServiceReadyCompleter.future;
Future<dynamic> get networkReady => networkServiceReadyCompleter.future;
Future<dynamic> get userServiceReady => userServiceReadyCompleter.future;
Future<dynamic> get authUserReady => userServiceReadyCompleter.future;
Future<dynamic> get cacheReady => cacheServiceReadyCompleter.future;
