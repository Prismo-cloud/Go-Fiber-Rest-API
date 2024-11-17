# Go REST API

Bu proje, Go ile yazılmış basit bir REST API'dir. API, farklı endpoint'ler ve middleware'ler ile temel bir yapı sunar. İlk defa bir REST API geliştirdiğim için bu proje, öğrenme sürecimdeki ilk adımları temsil ediyor.

## Özellikler
- **3 farklı endpoint:**
  - `GET /`: Basit bir karşılama mesajı döner.
  - `GET /user/:userId`: Belirtilen `userId` değerini döner.
  - `POST /user`: Kullanıcı bilgilerini bir modele göre alır ve doğrulama yapar.
- **4 middleware:**
  - **Logger Middleware**: `BaseURL`, `RequestURI` ve HTTP yöntemi (`Header Method`) bilgilerini döner.
  - **CorrelationId Middleware**: `/user` endpoint'inde `CorrelationId` kontrolü yapar.
  - **Recover Middleware**: Yapı panic olduğunda uygulamayı ayakta tutar.
  - **Panic Test Middleware**: Yapıyı panic ettirerek `Recover Middleware`'in çalışıp çalışmadığını kontrol eder.

## Kullanım

### Kurulum
Bu projeyi bilgisayarınıza klonlayın:
```bash
git clone https://github.com/username/repo-name.git
cd repo-name

Gerekli bağımlılıkları yüklemek için:

go mod tidy

Çalıştırma

API'yi başlatmak için:

go run main.go

Endpoint'ler
Metod	Endpoint	Açıklama
GET	/	Hello first get endpoint mesajını döner.
GET	/user/:userId	Verilen userId değerini döner.
POST	/user	Kullanıcı bilgisi alır ve doğrulama yapar.
Middleware'ler

    Logger Middleware
    Gelen her isteğin BaseURL, RequestURI ve Method bilgilerini konsola yazdırır.

    CorrelationId Middleware
    /user endpoint'ine gelen isteklerde CorrelationId başlığının olup olmadığını kontrol eder. Eğer eksikse, 400 hata kodu döner.

    Recover Middleware
    Panic durumunda uygulamayı ayakta tutar ve hata detaylarını loglar.

    Panic Test Middleware
    Yapıyı kasıtlı olarak panic ettirir ve Recover Middleware'in düzgün çalıştığını doğrular.

Geliştirme Notları

    İlk defa bir REST API geliştirdiğim için, geri bildirimlere açık bir yapıya sahibim. Projenin eksik veya geliştirilebilir yönlerini paylaşabilirsiniz.
    Recover Middleware özellikle hata durumlarının nasıl ele alındığını göstermek için eklendi.

Katkıda Bulunma

Katkıda bulunmak isterseniz bir fork oluşturun, değişikliklerinizi yapın ve bir pull request açın.
Lisans

Bu proje MIT lisansı ile lisanslanmıştır. Detaylar için LICENSE dosyasına bakabilirsiniz.

İyi kodlamalar! 😊