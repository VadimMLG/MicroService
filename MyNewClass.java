public class MyNewClass {
    public static void main(String[] args) {
        System.out.println("Мой новый класс в MicroService!");
        printMessage(); // Вызов нового метода
    }
    
    // Добавляем новый метод
    public static void printMessage() {
        System.out.println("Это новое изменение в классе!");
    }
}